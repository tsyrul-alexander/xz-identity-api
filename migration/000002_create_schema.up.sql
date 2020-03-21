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

CREATE TABLE "Role" (
    "Id" UUID PRIMARY KEY,
    "Name" VARCHAR (250),
    "Code" INTEGER NOT NULL
);

CREATE TABLE "UserRole" (
    "UserId" UUID REFERENCES "User"("Id"),
    "RoleId" UUID REFERENCES "Role"("Id"),
    PRIMARY KEY ("UserId", "RoleId")
);