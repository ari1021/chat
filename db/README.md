# データベース構成（MySQL）

- pr: primary key, 該当するものには o を書く
- uq: unique key, 該当するものには数字を書く、同じ数字は複合 unique key
- fk: foreign key, テーブルを書く

## rooms

| name        | Type     | pr  | required | uq  | fk(onupdate, ondelete) | default                     | description   |
| ----------- | -------- | --- | -------- | --- | ---------------------- | --------------------------- | ------------- |
| id          | int      | o   | True     |     |                        |                             | autoincrement |
| createdAt   | Datetime |     | True     |     |                        | current_timestamp           |
| updatedAt   | Datetime |     | True     |     |                        | current_timestamp on update |
| deletedAt   | Datetime |     |          |     |                        |                             | ログ保管のため  |
| name        | string   |     | True     | 1   |

## chats

| name      | Type     | pr  | required | uq  | fk(onupdate, ondelete)  | default                     | description  |
| --------- | -------- | --- | -------- | --- | ----------------------- | --------------------------- | ------------ |
| id        | int      | o   | True     |
| createdAt | Datetime |     | True     |     |                         | current_timestamp           |
| room_id   | int      |     | True     |     | rooms(cascade, cascade) |
| username  | string   |     | True     |
| chat      | string   |     | True     |
