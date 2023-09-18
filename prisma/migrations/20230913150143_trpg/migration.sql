-- CreateTable
CREATE TABLE "CharactorStatus" (
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

    CONSTRAINT "CharactorStatus_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "CharactorStatus_guildId_userId_key" ON "CharactorStatus"("guildId", "userId");
