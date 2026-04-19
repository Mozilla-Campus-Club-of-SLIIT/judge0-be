CREATE OR REPLACE FUNCTION update_submission_status()
RETURNS TRIGGER AS $$
DECLARE
    submission_row public.dsa_submissions%ROWTYPE;
    challenge_marks numeric;
    already_passed bigint;
BEGIN
    IF NEW.status = 2 THEN
        UPDATE public.dsa_submissions
        SET pass_count = pass_count + 1
        WHERE submission_id = NEW.submission_id;
    ELSIF NEW.status = 3 THEN
        UPDATE public.dsa_submissions
        SET fail_count = fail_count + 1
        WHERE submission_id = NEW.submission_id;
    END IF;
    
    SELECT * INTO submission_row
    FROM public.dsa_submissions
    WHERE submission_id = NEW.submission_id;

    IF submission_row.fail_count > 0 THEN
        UPDATE public.dsa_submissions
        SET evaluation_status = 3
        WHERE submission_id = NEW.submission_id;
    ELSIF submission_row.pass_count = submission_row.test_count THEN
        UPDATE public.dsa_submissions
        SET evaluation_status = 2
        WHERE submission_id = NEW.submission_id;

        SELECT marks INTO challenge_marks
        FROM public.challenges
        WHERE id = submission_row.challenge_id;

        SELECT COUNT(*) INTO already_passed
        FROM public.dsa_submissions
        WHERE user_id = submission_row.user_id
          AND challenge_id = submission_row.challenge_id
          AND evaluation_status = 2
          AND submission_id <> submission_row.submission_id;

        IF already_passed = 0 THEN
            INSERT INTO public.leaderboard(user_id, marks, last_updates)
            VALUES (submission_row.user_id, challenge_marks, now())
            ON CONFLICT (user_id)
            DO UPDATE SET 
                marks = leaderboard.marks + EXCLUDED.marks,
                last_updates = now();
        END IF;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_update_submission
AFTER INSERT ON public.dsa_submission_results
FOR EACH ROW
EXECUTE FUNCTION update_submission_status();