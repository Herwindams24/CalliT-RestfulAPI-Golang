-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 24 Bulan Mei 2021 pada 15.57
-- Versi server: 10.4.18-MariaDB
-- Versi PHP: 8.0.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `goapidb`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `daftarinstansi`
--

CREATE TABLE `daftarinstansi` (
  `id` int(11) NOT NULL,
  `namaInstansi` varchar(100) NOT NULL,
  `kota` varchar(50) NOT NULL,
  `telfon` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `daftarinstansi`
--

INSERT INTO `daftarinstansi` (`id`, `namaInstansi`, `kota`, `telfon`) VALUES
(1, 'RSUD dr. Soegiri Lamongan', 'Lamongan', '322321718'),
(2, 'Rumah Sakit Muhammadiyah Lamongan', 'Lamongan', '322322834'),
(3, 'Rumah Sakit Citra Medika Lamongan', 'Lamongan', '322312444'),
(4, 'RSI Nashrul Ummah', 'Lamongan', '322321522'),
(5, 'Rumah Sakit Husada Utama', 'Surabaya', '315018335'),
(6, 'Rumah Sakit Darmo', 'Surabaya', '315676253'),
(7, 'RSUD Dr. Soetomo', 'Surabaya', '315501078'),
(8, 'RS PHC Surabaya', 'Surabaya', '313294801'),
(9, 'Rumah Sakit Semen Gresik', 'Gresik', '313987840'),
(10, 'RS. Muhammadiyah Gresik', 'Gresik', '313981275'),
(11, 'Rumah Sakit Umum Rachmi Dewi', 'Gresik', '313957448'),
(12, 'RSUD Ibnu Sina', 'Gresik', '313951239'),
(13, 'Rumah Sakit Aisyiyah Bojonegoro', 'Bojonegoro', '353881748'),
(14, 'RSUD Dr.Sosodoro Bojonegoro', 'Bojonegoro', '353341213'),
(15, 'IGD RS Bhayangkara Wahyu Tutuko Bojonegoro', 'Bojonegoro', '353888780'),
(16, 'Rumah Sakit Islam Pemuda', 'Bojonegoro', '353887749'),
(20, 'Rumah Sakit Umum Daerah Sidoarjo', 'Lamongan', '0318961649');

-- --------------------------------------------------------

--
-- Struktur dari tabel `daftarkota`
--

CREATE TABLE `daftarkota` (
  `id_kota` int(11) NOT NULL,
  `nama_kota` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `daftarkota`
--

INSERT INTO `daftarkota` (`id_kota`, `nama_kota`) VALUES
(1, 'Lamongan'),
(2, 'Surabaya'),
(3, 'Gresik'),
(4, 'Bojonegoro');

-- --------------------------------------------------------

--
-- Struktur dari tabel `jenisinstansi`
--

CREATE TABLE `jenisinstansi` (
  `id_instansi` int(11) DEFAULT NULL,
  `jenis_instansi` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `jenisinstansi`
--

INSERT INTO `jenisinstansi` (`id_instansi`, `jenis_instansi`) VALUES
(1, 'Rumah Sakit'),
(2, 'Kantor Polisi'),
(1, 'Rumah Sakit'),
(2, 'Kantor Polisi');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `username`, `password`) VALUES
(1, 'Herwinda24', '$2a$10$aXitB3Uu4oYRgUmB8iWlBOXn4geS5VacRTz7rZ4FkXGRdyIoiudvC');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `daftarinstansi`
--
ALTER TABLE `daftarinstansi`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `daftarkota`
--
ALTER TABLE `daftarkota`
  ADD PRIMARY KEY (`id_kota`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `daftarinstansi`
--
ALTER TABLE `daftarinstansi`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=28;

--
-- AUTO_INCREMENT untuk tabel `daftarkota`
--
ALTER TABLE `daftarkota`
  MODIFY `id_kota` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
