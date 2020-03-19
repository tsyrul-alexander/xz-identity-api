CREATE TABLE "DefaultIdentity" (
    "Id" UUID PRIMARY KEY,
    "Login" VARCHAR (250),
    "Password" VARCHAR (250)
);

CREATE TABLE "User" (
    "Id" UUID PRIMARY KEY,
    "Name" VARCHAR (250),
    "IdentityType" INTEGER NOT NULL,
    "DefaultIdentityId" UUID REFERENCES "DefaultIdentity"("Id")
);