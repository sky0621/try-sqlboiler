
-- +migrate Up
INSERT INTO customer(name, age, prefectures) VALUES
  ('Sato', 30, 'Tokyo'),
  ('Suzuki', 22, 'Osaka'),
  ('Takahashi', 27, 'Yamagata'),
  ('Tanaka', 19, 'Nagasaki'),
  ('Ito', 14, 'Sizuoka'),
  ('Watanabe', 33, 'Okinawa'),
  ('Yamamoto', 25, 'Tokyo'),
  ('Nakamura', 29, 'Saitama'),
  ('Kobayashi', 17, 'Gifu'),
  ('Kato', 20, 'Miyagi'),
  ('Yoshida', 35, 'Hokkaido'),
  ('Yamada', 20, 'Hokkaido'),
  ('Sasaki', 24, 'Tokyo'),
  ('Yamaguchi', 22, 'Ehime'),
  ('Matsumoto', 38, 'Tokyo');
-- +migrate Down
