-- CreateTable
CREATE TABLE "Point" (
    "id" UUID NOT NULL,
    "userId" VARCHAR(19) NOT NULL,
    "guildId" VARCHAR(19) NOT NULL,
    "value" INTEGER NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "Point_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "Point_userId_guildId_key" ON "Point"("userId", "guildId");
