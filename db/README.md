# データベース構成（MySQL）

- pr: primary key, 該当するものには o を書く
- uq: unique key, 該当するものには数字を書く、同じ数字は複合 unique key
- fk: foreign key, テーブルを書く

## rooms

| name        | Type     | pr  | required | uq  | fk(onupdate, ondelete)  | default                     | description   |
| ----------- | -------- | --- | -------- | --- | ----------------------- | --------------------------- | ------------- |
| id          | int      | o   | True     |     |                         |                             | autoincrement |
| createdAt   | Datetime |     | True     |     |                         | current_timestamp           |
| updatedAt   | Datetime |     | True     |     |                         | current_timestamp on update |
| deletedAt   | Datetime |     |          |     |                         |                             | ログ保管のため  |
| name        | string   |     | True     | 1   |
| user_id     | int      |     | True     |     | users(cascade, cascade) |


## chats

| name      | Type     | pr  | required | uq  | fk(onupdate, ondelete)  | default                     | description   |
| --------- | -------- | --- | -------- | --- | ----------------------- | --------------------------- | ------------- |
| id        | int      | o   | True     |     |                         |                             | autoincrement |
| createdAt | Datetime |     | True     |     |                         | current_timestamp           |
| room_id   | int      |     | True     |     | rooms(cascade, cascade) |
| user_id   | string   |     | True     |     | users(cascade, cascade) |
| message   | string   |     | True     |


## users

| name               | Type     | pr  | required | uq  | fk(onupdate, ondelete)  | default                     | description   |
| ------------------ | -------- | --- | -------- | --- | ----------------------- | --------------------------- | ------------- |
| id                 | int      | o   | True     |     |                         |                             | autoincrement |
| createdAt          | Datetime |     | True     |     |                         | current_timestamp           |
| updatedAt          | Datetime |     | True     |     |                         | current_timestamp on update |
| deletedAt          | Datetime |     |          |     |                         |                             | ログ保管のため  |
| name               | string   |     | True     |
| id_token           | string   |     |          |     |                         |                             | 1000文字以上   |
| access_token       | string   |     |          |     |                         |                             | 1000文字以上   |
| refresh_token      | string   |     |          |     |                         |                             | 1000文字以上   |