DROP SCHEMA imageRepo;
CREATE SCHEMA imageRepo;
Use imageRepo;
	CREATE TABLE imageRepo.images (
        image_id INT NOT NULL AUTO_INCREMENT,
        path VARCHAR(255) NOT NULL,
        PRIMARY KEY (image_id),
        UNIQUE INDEX idnew_table_UNIQUE (image_id ASC),
        UNIQUE INDEX path_UNIQUE (path ASC)
    );

    CREATE TABLE users (
        user_id int(11) NOT NULL AUTO_INCREMENT,
        username varchar(50) NOT NULL,
        PRIMARY KEY (user_id),
        UNIQUE KEY userscol_UNIQUE (username),
        UNIQUE KEY user_id_UNIQUE (user_id)
    );

    CREATE TABLE users_images (
        ui_id int(11) NOT NULL AUTO_INCREMENT,
        user_id int(11) NOT NULL,
        image_id int(11) NOT NULL,
        PRIMARY KEY (ui_id),
        FOREIGN KEY (user_id) REFERENCES users(user_id),
        FOREIGN KEY (image_id) REFERENCES images(image_id),
        UNIQUE KEY ui_id_UNIQUE (ui_id)
    ); 
