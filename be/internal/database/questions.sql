drop table if exists public.questions;
create table if not exists public.questions(
    id uuid primary key default gen_random_uuid(),
    question text unique,
    answer text,
    answered_at date,
    created_at date
);