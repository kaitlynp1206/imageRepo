
CREATE SCHEMA imageRepo;
	CREATE TABLE images (
		image_id INT AUTO_INCREMENT PRIMARY KEY,
        path VARCHAR(255) NOT NULL,
        description VARCHAR(50),
        UNIQUE KEY unique_path (path)
	);

    CREATE TABLE users (
        user_id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(50) NOT NULL,
        password VARCHAR(255) NOT NULL,
        UNIQUE KEY unique_username (username)
    );

    CREATE TABLE users_images (
        ui_id INT AUTO_INCREMENT PRIMARY KEY,
        user_id INT NOT NULL REFERENCES users(user_id),
        image_id INT NOT NULL REFERENCES images(image_id)
    );
