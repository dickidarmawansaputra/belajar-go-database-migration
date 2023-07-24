CREATE TABLE correct(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

CREATE TABLE wrong(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

/*
Dirty State = kondisi dimana tidak bisa melakukan up / down migration. ini terjadi karna ada kesalahan syntax di file migration
Solusi:
1. setelah diperbaiki manual syntaxnya, perlu mengubah versi migration di table schema_migrations
2. check versi jika dirty = 1 maka update ke versi sebelumnya di table schema_migrations 
migrate -database "mysql://dickids:rahasia@tcp(localhost:3306)/belajar-go-database-migration" -path database/migrations version
3. untuk update versi dengan command ini:
migrate -database "mysql://dickids:rahasia@tcp(localhost:3306)/belajar-go-database-migration" -path database/migrations force 20230721064031
4. check versinya lagi dan pastikan dirty = 0
5. drop table correct
6. migrate up lagi
*/
