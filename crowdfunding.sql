-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 25 Mar 2021 pada 19.10
-- Versi server: 10.4.13-MariaDB
-- Versi PHP: 7.2.32

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `crowdfunding`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `campaigns`
--

CREATE TABLE `campaigns` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `name` varchar(128) NOT NULL,
  `short_description` varchar(128) NOT NULL,
  `description` varchar(256) NOT NULL,
  `goal_amount` int(11) NOT NULL,
  `current_amount` int(11) NOT NULL,
  `perks` varchar(128) NOT NULL,
  `baker_count` int(11) NOT NULL,
  `slug` varchar(128) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `campaigns`
--

INSERT INTO `campaigns` (`id`, `user_id`, `name`, `short_description`, `description`, `goal_amount`, `current_amount`, `perks`, `baker_count`, `slug`, `created_at`, `updated_at`) VALUES
(1, 1, 'Game Jabari Parker', 'ini adlaha game jabari parker', 'dasddddddddddddddddddddd jabarikdsd asjdasdasdd', 1000000, 0, 'mendapat uang,mendapat kehormatan,mendapat wanita', 0, 'game-jabari-parker', '2021-03-14 02:04:12', '2021-03-14 02:04:12'),
(2, 1, 'Game jabari 2', 'sdadasd ini adlaah desc', 'sdadasd ini adlaah descsdadasd ini adlaah descsdadasd ini adlaah desc', 100000000, 0, 'harta,tahta,wanita', 0, 'game-jabari-2', '2021-03-14 02:04:12', '2021-03-14 02:04:12'),
(3, 2, 'Game Haikyu', 'ini adalah game Haikyu', 'ini adalah game Haikyu ini adalah game Haikyu ini adalah game Haikyu', 200000000, 0, 'mendapat tanda tangan oikawa,tanda tangan ushijima,jersey kageyama', 0, 'game-haikyu', '2021-03-14 02:06:37', '2021-03-14 02:06:37'),
(7, 0, 'Pembangunan rumah nenek', '', '', 0, 0, '', 0, '', '2021-03-21 04:38:33', '0000-00-00 00:00:00'),
(8, 0, 'Pembangunan rumah kakek', 'Short', 'LOooongg', 100000000, 0, 'Ganteng Iyak', 0, '1-Pembangunan-rumah-kakek', '2021-03-21 05:00:11', '0000-00-00 00:00:00'),
(9, 0, 'Pembangunan rumah kakek', 'Short', 'LOooongg', 100000000, 0, 'Ganteng Iyak', 0, '0-Pembangunan-rumah-kakek', '2021-03-25 02:40:16', '0000-00-00 00:00:00'),
(10, 16, 'Pembangunan rumah kakek', 'Short', 'LOooongg', 100000000, 0, 'Ganteng Iyak', 0, '16-Pembangunan-rumah-kakek', '2021-03-25 02:44:02', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Struktur dari tabel `campaign_images`
--

CREATE TABLE `campaign_images` (
  `id` int(11) NOT NULL,
  `campaign_id` int(11) NOT NULL,
  `file_name` varchar(128) NOT NULL,
  `is_primary` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `campaign_images`
--

INSERT INTO `campaign_images` (`id`, `campaign_id`, `file_name`, `is_primary`, `created_at`, `updated_at`) VALUES
(1, 1, 'image1.png', 1, '2021-03-14 02:32:16', '2021-03-14 02:32:16'),
(2, 1, 'image2.jpg', 0, '2021-03-14 02:32:16', '2021-03-14 02:32:16'),
(3, 2, 'gamabr2.png', 1, '2021-03-14 02:33:01', '2021-03-14 02:33:01');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(128) DEFAULT NULL,
  `occupation` varchar(128) DEFAULT NULL,
  `email` varchar(128) NOT NULL,
  `password` varchar(256) NOT NULL,
  `image` varchar(128) DEFAULT NULL,
  `role` varchar(64) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password`, `image`, `role`, `created_at`, `updated_at`) VALUES
(16, 'Harits Rizkal Bokuto', 'Pelajar', 'ranjau010401@gmail.com', '$2a$04$8hB45.ndZLtYOq/DtU4vN.SIPXCOT9NlSBK7HnUDSDDSpuF/oiYvG', 'images/16-gudeg bromo.jpg', 'user', '2021-03-05 15:08:05', '2021-03-13 15:45:18'),
(17, 'Harits Rizkal', 'Pelajar', 'rizkal@gmail.com', '$2a$04$0fyiE8pTRydznPh1jTbAiurjqmUskh6pc3WqzVAsexrMjblMsJywa', '', 'user', '2021-03-05 15:23:46', '2021-03-05 15:23:46'),
(18, 'Kageyama Tobio', 'Volleyball Player', 'kageyama@gmail.com', '$2a$04$0mEOBWT8qJmLBS9hdNcAo.LG6eJzTeuGTdpu.RhT/wGQCInTpFeMK', 'images/18-gudeg bromo.jpg', 'user', '2021-03-12 19:43:43', '2021-03-13 16:01:12'),
(19, 'Hinata Shoyo', 'Volleyball Player', 'hinata@gmail.com', '$2a$04$mvyBGgT4dWR5Bts2.xtrVOYdtDZ4iBP7y8hs.GwI0Isk3o69LyoOu', '', 'user', '2021-03-12 19:51:36', '2021-03-12 19:51:36'),
(20, 'Oikawa', 'Volleyball Player', 'oikawa@gmail.com', '$2a$04$ZA9HWRWuC5Xy8okY042q6uIaZ4gohRfrMvdPvG7CAnvH/Js1MgXiW', '', 'user', '2021-03-12 19:52:45', '2021-03-12 19:52:45'),
(21, 'Ushijima', 'Volleyball Player', 'ushijima@gmail.com', '$2a$04$iD68S4eIPhni2IKgcy9qkO9pboH9VsO348CNzg2exP2LwgtFbAani', 'images/21-gudeg bromo.jpg', 'user', '2021-03-12 19:53:37', '2021-03-13 15:57:17');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `campaigns`
--
ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `campaign_images`
--
ALTER TABLE `campaign_images`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT untuk tabel `campaign_images`
--
ALTER TABLE `campaign_images`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
