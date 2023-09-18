/*
  Warnings:

  - You are about to drop the `CharactorStatus` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropTable
DROP TABLE "CharactorStatus";

-- CreateTable
CREATE TABLE "CharacterStatus" (
    "id" UUID NOT NULL,
    "guildId" VARCHAR(19) NOT NULL,
    "userId" VARCHAR(19) NOT NULL,
    "str" INTEGER NOT NULL,
    "con" INTEGER NOT NULL,
    "pow" INTEGER NOT NULL,
    "dex" INTEGER NOT NULL,
    "app" INTEGER NOT NULL,
    "siz" INTEGER NOT NULL,
    "int" INTEGER NOT NULL,
    "edu" INTEGER NOT NULL,
    "luk" INTEGER NOT NULL,
    "san" INTEGER NOT NULL,
    "ida" INTEGER NOT NULL,
    "know" INTEGER NOT NULL,
    "hp" INTEGER NOT NULL,
    "mov" INTEGER NOT NULL,
    "mp" INTEGER NOT NULL,
    "skl" INTEGER NOT NULL,
    "db" INTEGER NOT NULL,
    "build" INTEGER NOT NULL,

    CONSTRAINT "CharacterStatus_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "CharacterStatus_guildId_userId_key" ON "CharacterStatus"("guildId", "userId");
