CREATE OR REPLACE VIEW get_dsa_challenges_view AS
SELECT
  c.id,
  c.created_at,
  c.title,
  c.description,
  c.type_id,
  c.status_id,
  c.marks,
  ct.type AS type,
  cs.status AS status,
  dc.sample_input,
  dc.sample_output,
  dc.note
FROM public.challenges c
JOIN public.challenge_types ct ON c.type_id = ct.id
JOIN public.challenge_status cs ON c.status_id = cs.id
LEFT JOIN public.dsa_challenges dc ON c.id = dc.challenge_id
WHERE c.type_id = 1;


CREATE OR REPLACE VIEW preview_challenges_view AS
SELECT
  c.id,
  c.created_at,
  c.title,
  c.description,
  c.type_id,
  c.status_id,
  ct.type AS type,
  cs.status AS status,
  c.marks
FROM public.challenges c
JOIN public.challenge_types ct ON c.type_id = ct.id
JOIN public.challenge_status cs ON c.status_id = cs.id;