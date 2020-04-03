
CREATE TABLE IF NOT EXISTS alteration_method (
    id smallint PRIMARY KEY,
    default_name text NOT NULL,
    is_deleted boolean DEFAULT false NOT NULL,
    src_time timestamp(3) with time zone NOT NULL,
    imported_at timestamp(3) with time zone DEFAULT CURRENT_TIMESTAMP(3) NOT NULL
);

INSERT INTO alteration_method (id, default_name, is_deleted, src_time, imported_at) VALUES (1, 'blind_stitch_order_made_jacket', false, '2010-06-01 23:59:59.371+00', '2019-08-16 02:20:28.505+00');
INSERT INTO alteration_method (id, default_name, is_deleted, src_time, imported_at) VALUES (2, 'single_stitch', false, '2010-06-01 23:59:59.371+00', '2019-08-16 02:20:28.505+00');
INSERT INTO alteration_method (id, default_name, is_deleted, src_time, imported_at) VALUES (3, 'single_stitch_shirt_sleeve', false, '2010-06-01 23:59:59.371+00', '2019-08-16 02:20:28.505+00');
