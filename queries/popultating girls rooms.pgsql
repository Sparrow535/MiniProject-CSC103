WITH blocks AS (
  SELECT unnest(ARRAY['O', 'P']) AS block
),
floors AS (
  SELECT unnest(ARRAY['F', 'S', 'T', 'Fo']) AS floor
),
room_numbers AS (
  SELECT
    block,
    floor,
    block || floor || '-' || generate_series(1, 8) AS room_no
  FROM blocks
  CROSS JOIN floors
)
INSERT INTO roomg (block, floor, room_no)
SELECT block, floor, room_no
FROM room_numbers;
