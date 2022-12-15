-- SQLite
--CREATE TABLE IF NOT EXISTS User (
--	UserID int PRIMARY KEY,
--  	username string NOT NULL
	
--) WITHOUT ROWID;


--INSERT into Users(UserID, username) values (1, "valerio");

--DROP TABLE Banned;

--CREATE TABLE IF NOT EXISTS Follower (
--	FollowerID int ,
--    FollowedID int,
--    PRIMARY KEY(FollowerID, FOllowedID)
--	FOREIGN KEY (FollowerID) 
--      REFERENCES Users (UserID) 
--         ON DELETE CASCADE 
--         ON UPDATE NO ACTION
--    FOREIGN KEY (FollowedID) 
--      REFERENCES Users (UserID) 
--         ON DELETE CASCADE 
--         ON UPDATE NO ACTION
--) WITHOUT ROWID;

CREATE TABLE IF NOT EXISTS Banned (
	BannedID int,
    BanningID int,
    PRIMARY KEY (BannedID, BanningID)
	FOREIGN KEY (BannedID) 
      REFERENCES Users (UserID) 
        ON DELETE CASCADE 
        ON UPDATE NO ACTION
    FOREIGN KEY (BanningID) 
      REFERENCES Users (UserID) 
         ON DELETE CASCADE 
         ON UPDATE NO ACTION
) WITHOUT ROWID;


--CREATE TABLE IF NOT EXISTS Photo (
--	PhotoID int,
--    UserID int,
--    Photo string NOT NULL,
--    DataPost string NOT NULL,
--	PRIMARY KEY (PhotoID, UserID)
--    FOREIGN KEY (UserID) 
--      REFERENCES Users (UserID) 
--        ON DELETE CASCADE 
--         ON UPDATE NO ACTION
--) WITHOUT ROWID;

--CREATE TABLE IF NOT EXISTS Comment (
--	PhotoID int,
--    UserIDReceiving int NOT NULL,
--    CommentID int,
--    UserIDSending int NOT NULL,
--    DataPost string NOT NULL,
--	PRIMARY KEY (PhotoID, CommentID)
--    FOREIGN KEY (UserIDReceiving) 
--      REFERENCES Users (UserID) 
--       ON DELETE CASCADE 
--         ON UPDATE NO ACTION
--    FOREIGN KEY (UserIDSending) 
--      REFERENCES Users (UserID) 
--       ON DELETE CASCADE 
--         ON UPDATE NO ACTION
--    FOREIGN KEY (PhotoID) 
--      REFERENCES Users (UserID) 
--       ON DELETE CASCADE 
--         ON UPDATE NO ACTION
--) WITHOUT ROWID;

 -- CREATE TABLE IF NOT EXISTS Like (
 
 --     UserIDPutting int,
 --     PhotoID string NOT NULL,
 --     UserIDReceiving int NOT NULL,
 --     DataPost string NOT NULL,
 --     PRIMARY KEY (PhotoID, UserIDPutting)
 --     FOREIGN KEY (UserIDReceiving) 
 --       REFERENCES Users (UserID) 
--          ON DELETE CASCADE 
--           ON UPDATE NO ACTION
--      FOREIGN KEY (UserIDPutting) 
--        REFERENCES Users (UserID) 
--          ON DELETE CASCADE 
--           ON UPDATE NO ACTION
--      FOREIGN KEY (PhotoID) 
--        REFERENCES Photo (PhotoID) 
--          ON DELETE CASCADE 
--           ON UPDATE NO ACTION
--  ) WITHOUT ROWID;

--ALTER TABLE Like DROP COLUMN LikeID;