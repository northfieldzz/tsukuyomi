-- CreateTable
CREATE TABLE "NotificationChannel" (
    "id" UUID NOT NULL,
    "guildId" VARCHAR(19) NOT NULL,
    "channelId" VARCHAR(19) NOT NULL,

    CONSTRAINT "NotificationChannel_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "NotificationChannel_guildId_key" ON "NotificationChannel"("guildId");
