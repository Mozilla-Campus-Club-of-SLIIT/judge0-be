-- WARNING: This schema is for context only and is not meant to be run.
-- Table order and constraints may not be valid for execution.

CREATE TABLE public.challenge_status (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  status text NOT NULL,
  CONSTRAINT challenge_status_pkey PRIMARY KEY (id)
);
CREATE TABLE public.challenge_types (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  type text NOT NULL DEFAULT ''::text UNIQUE,
  CONSTRAINT challenge_types_pkey PRIMARY KEY (id)
);
CREATE TABLE public.challenges (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  title text NOT NULL,
  description text NOT NULL,
  type_id bigint NOT NULL,
  status_id bigint NOT NULL,
  marks numeric NOT NULL DEFAULT '5'::numeric,
  CONSTRAINT challenges_pkey PRIMARY KEY (id),
  CONSTRAINT challenges_type_id_fkey FOREIGN KEY (type_id) REFERENCES public.challenge_types(id),
  CONSTRAINT challenges_status_id_fkey FOREIGN KEY (status_id) REFERENCES public.challenge_status(id)
);
CREATE TABLE public.dsa_challenges (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  sample_input text NOT NULL,
  sample_output text,
  challenge_id bigint NOT NULL UNIQUE,
  note text,
  CONSTRAINT dsa_challenges_pkey PRIMARY KEY (id),
  CONSTRAINT dsa_challenges_challenge_id_fkey FOREIGN KEY (challenge_id) REFERENCES public.challenges(id)
);
CREATE TABLE public.dsa_result_status (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  status_id numeric NOT NULL UNIQUE,
  status text NOT NULL UNIQUE,
  CONSTRAINT dsa_result_status_pkey PRIMARY KEY (id)
);
CREATE TABLE public.dsa_submission_results (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  submission_id text NOT NULL,
  status numeric NOT NULL DEFAULT '1'::numeric,
  token text NOT NULL UNIQUE,
  CONSTRAINT dsa_submission_results_pkey PRIMARY KEY (id),
  CONSTRAINT dsa_submission_results_status_fkey FOREIGN KEY (status) REFERENCES public.dsa_result_status(status_id),
  CONSTRAINT dsa_submission_results_submission_id_fkey FOREIGN KEY (submission_id) REFERENCES public.dsa_submissions(submission_id)
);
CREATE TABLE public.dsa_submissions (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  submission_id text NOT NULL UNIQUE,
  challenge_id bigint NOT NULL,
  user_id text NOT NULL,
  test_count bigint NOT NULL,
  pass_count bigint NOT NULL DEFAULT '0'::bigint,
  fail_count bigint NOT NULL DEFAULT '0'::bigint,
  evaluation_status numeric NOT NULL DEFAULT '1'::numeric,
  CONSTRAINT dsa_submissions_pkey PRIMARY KEY (id),
  CONSTRAINT dsa_submissions_challenge_id_fkey FOREIGN KEY (challenge_id) REFERENCES public.challenges(id)
);
CREATE TABLE public.dsa_test_cases (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  challenge_id bigint NOT NULL,
  test_input text DEFAULT ' '::text,
  test_output text DEFAULT ' '::text,
  CONSTRAINT dsa_test_cases_pkey PRIMARY KEY (id),
  CONSTRAINT dsa_test_cases_challenge_id_fkey FOREIGN KEY (challenge_id) REFERENCES public.challenges(id)
);
CREATE TABLE public.evaluation_status (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  status_id numeric NOT NULL,
  status text NOT NULL,
  CONSTRAINT evaluation_status_pkey PRIMARY KEY (id)
);
CREATE TABLE public.leaderboard (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  user_id text NOT NULL UNIQUE,
  marks numeric NOT NULL DEFAULT '0'::numeric,
  last_updates timestamp with time zone NOT NULL DEFAULT now(),
  CONSTRAINT leaderboard_pkey PRIMARY KEY (id),
  CONSTRAINT leaderboard_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id)
);
CREATE TABLE public.settings (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  setting text,
  value boolean,
  CONSTRAINT settings_pkey PRIMARY KEY (id)
);
CREATE TABLE public.users (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  user_id text NOT NULL UNIQUE,
  email text NOT NULL UNIQUE,
  name text NOT NULL UNIQUE,
  CONSTRAINT users_pkey PRIMARY KEY (id)
);