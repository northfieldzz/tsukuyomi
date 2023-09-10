/*
  Warnings:

  - The primary key for the `Invite` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `expireAt` on the `Invite` table. All the data in the column will be lost.
  - You are about to drop the column `id` on the `Invite` table. All the data in the column will be lost.
  - You are about to drop the column `madeBy` on the `Invite` table. All the data in the column will be lost.
  - You are about to drop the column `useCount` on the `Invite` table. All the data in the column will be lost.

*/
-- DropIndex
DROP INDEX "Invite_madeBy_code_key";

-- AlterTable
ALTER TABLE "Invite" DROP CONSTRAINT "Invite_pkey",
DROP COLUMN "expireAt",
DROP COLUMN "id",
DROP COLUMN "madeBy",
DROP COLUMN "useCount",
ADD COLUMN     "expires" TIMESTAMP(3),
ADD COLUMN     "inviterId" VARCHAR(19),
ADD COLUMN     "uses" INTEGER NOT NULL DEFAULT 0,
ADD CONSTRAINT "Invite_pkey" PRIMARY KEY ("code");
