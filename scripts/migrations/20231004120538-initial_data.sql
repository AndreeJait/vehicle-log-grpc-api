
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
INSERT INTO public.types (id, name) VALUES
    ('923d6f04-9bc9-4c0d-b3e1-c2e60897f376', 'Car'),
    ('fec2f4f1-0f7c-4cbe-90e2-d3b22d7c729b', 'Bus'),
    ('1f1ac6f4-6aea-4255-bc62-648ee88586f1', 'Mini Bus'),
    ('796eb60f-9752-4c1d-821a-9bab3a4dbbc0', 'Motorcycle');
INSERT INTO public.towns (id, town_code, town_name) VALUES
    ('afd5e1a4-cdb2-484b-9a51-81d889bb4d68', 'T001', 'Tower Example 01'),
    ('eb7cba8a-4324-4ca0-b8d7-b643cadadc38', 'T002', 'Tower Example 02');
INSERT INTO public.town_slots (id, town_id, type_id, quantity, price) VALUES
    (uuid_generate_v4(), 'afd5e1a4-cdb2-484b-9a51-81d889bb4d68', '923d6f04-9bc9-4c0d-b3e1-c2e60897f376', 250, 2000),
    (uuid_generate_v4(), 'afd5e1a4-cdb2-484b-9a51-81d889bb4d68', 'fec2f4f1-0f7c-4cbe-90e2-d3b22d7c729b', 100, 3000),
    (uuid_generate_v4(), 'afd5e1a4-cdb2-484b-9a51-81d889bb4d68', '796eb60f-9752-4c1d-821a-9bab3a4dbbc0', 50, 5000),
    (uuid_generate_v4(), 'afd5e1a4-cdb2-484b-9a51-81d889bb4d68', '1f1ac6f4-6aea-4255-bc62-648ee88586f1', 50, 5000),
    (uuid_generate_v4(), 'eb7cba8a-4324-4ca0-b8d7-b643cadadc38', '923d6f04-9bc9-4c0d-b3e1-c2e60897f376', 150, 2000),
    (uuid_generate_v4(), 'eb7cba8a-4324-4ca0-b8d7-b643cadadc38', 'fec2f4f1-0f7c-4cbe-90e2-d3b22d7c729b', 100, 3000);
-- +migrate Down
