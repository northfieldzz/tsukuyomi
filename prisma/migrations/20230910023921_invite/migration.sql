-- CreateTable
CREATE TABLE "Invite" (
    "id" TEXT NOT NULL,
    "madeBy" VARCHAR(19) NOT NULL,
    "code" VARCHAR(16) NOT NULL,
    "useCount" INTEGER NOT NULL DEFAULT 0,
    "expireAt" TIMESTAMP(3),

    CONSTRAINT "Invite_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "Invite_madeBy_code_key" ON "Invite"("madeBy", "code");
