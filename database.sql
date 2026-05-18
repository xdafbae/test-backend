-- ============================================================
-- SharingVision 2023 — Post Article App
-- Database Setup Script
-- ============================================================

CREATE DATABASE IF NOT EXISTS article;
USE article;

CREATE TABLE IF NOT EXISTS posts (
    id           INT AUTO_INCREMENT PRIMARY KEY,
    title        VARCHAR(200)  NOT NULL,
    content      TEXT          NOT NULL,
    category     VARCHAR(100)  NOT NULL,
    created_date TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    status       VARCHAR(100)  NOT NULL
);

-- Sample data (optional)
INSERT INTO posts (title, content, category, status) VALUES
(
    'Pengenalan Golang untuk Pemula',
    'Golang atau Go adalah bahasa pemrograman open-source yang dikembangkan oleh Google. Bahasa ini dirancang untuk kesederhanaan, efisiensi, dan kemudahan dalam membangun sistem yang scalable. Dalam artikel ini, kita akan membahas dasar-dasar Golang mulai dari instalasi hingga membuat program pertama Anda. Go memiliki syntax yang bersih dan mudah dipahami, serta dilengkapi dengan garbage collector sehingga manajemen memori menjadi lebih mudah.',
    'Technology',
    'publish'
),
(
    'Membangun REST API dengan Gin Framework',
    'Gin adalah salah satu web framework paling populer di ekosistem Go. Framework ini terkenal karena performanya yang sangat cepat dan kemudahan penggunaannya. Dalam artikel ini, kita akan membangun sebuah REST API lengkap menggunakan Gin, mulai dari routing, middleware, validasi input, hingga koneksi ke database MySQL menggunakan GORM. Dengan mengikuti panduan ini, Anda akan mampu membuat API yang production-ready.',
    'Technology',
    'publish'
),
(
    'Tips Produktivitas untuk Developer',
    'Menjadi developer yang produktif bukan hanya tentang menulis kode dengan cepat, tetapi juga tentang menulis kode yang berkualitas dan mudah dipelihara. Dalam artikel ini, kami berbagi tips dan trik yang telah terbukti meningkatkan produktivitas developer, mulai dari penggunaan shortcut keyboard, manajemen waktu dengan teknik Pomodoro, pentingnya dokumentasi, hingga cara efektif melakukan code review.',
    'Productivity',
    'draft'
);
