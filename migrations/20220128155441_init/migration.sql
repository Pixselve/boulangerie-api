/*
  Warnings:

  - You are about to drop the `Post` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropTable
PRAGMA foreign_keys=off;
DROP TABLE "Post";
PRAGMA foreign_keys=on;

-- CreateTable
CREATE TABLE "Bakery" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL
);

-- CreateTable
CREATE TABLE "Bread" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "price" REAL NOT NULL,
    "bakeryID" TEXT NOT NULL,
    CONSTRAINT "Bread_bakeryID_fkey" FOREIGN KEY ("bakeryID") REFERENCES "Bakery" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
