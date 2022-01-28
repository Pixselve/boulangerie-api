/*
  Warnings:

  - A unique constraint covering the columns `[name]` on the table `Bakery` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "Bakery_name_key" ON "Bakery"("name");
