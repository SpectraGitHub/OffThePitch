--
-- Ny fysisk modell
--
CREATE DATABASE otp;

use otp;
DROP TABLE IF EXISTS bruker;
CREATE TABLE bruker (
    bruker_id INT NOT NULL AUTO_INCREMENT,
    brukernavn VARCHAR(128) NOT NULL,
    poeng INT NOT NULL,
    epost VARCHAR(256) NOT NULL,
    passord VARCHAR(128) NOT NULL,
    PRIMARY KEY (`bruker_id`),
    UNIQUE KEY unik_brukernavn(`brukernavn`),
    UNIQUE KEY unik_epost(`epost`)
);

DROP TABLE IF EXISTS spiller;
CREATE TABLE spiller (
    spiller_id INT NOT NULL,
    fornavn VARCHAR(128) NOT NULL,
    etternavn VARCHAR(128) NOT NULL,
    lag_id INT NOT NULL,
    maal INT NOT NULL,
    PRIMARY KEY(`spiller_id`)
);

DROP TABLE IF EXISTS toppscorer;
CREATE TABLE toppscorer (
    bruker_id INT NOT NULL,
    spiller_id INT NOT NULL,
    PRIMARY KEY(`bruker_id`, `spiller_id`),
    CONSTRAINT FK_ToppscorerBruker FOREIGN KEY (bruker_id) REFERENCES bruker(bruker_id)
   -- CONSTRAINT FK_ToppscorerSpiller FOREIGN KEY (spiller_id) REFERENCES spiller(spiller_id)
);

DROP TABLE IF EXISTS lagnavn;
CREATE TABLE lagnavn (
    lag_id INT NOT NULL,
    lag_navn VARCHAR(128) NOT NULL,
    PRIMARY KEY(`lag_id`)
);

DROP TABLE IF EXISTS kamper;
CREATE TABLE kamper (
    kamp_id INT NOT NULL,
    borte_id INT NOT NULL,
    borte_score INT NOT NULL,
    hjemme_id INT NOT NULL,
    hjemme_score INT NOT NULL,
    PRIMARY KEY(`kamp_id`),
    CONSTRAINT FK_Kamper_hjemme_lagnavn FOREIGN KEY (hjemme_id) REFERENCES lagnavn(lag_id),
    CONSTRAINT FK_Kamper_borte_lagnavn FOREIGN KEY (borte_id) REFERENCES lagnavn(lag_id)
);

DROP TABLE IF EXISTS kamp_valg;
CREATE TABLE kamp_valg (
    bruker_id INT NOT NULL,
    kamp_id INT NOT NULL,
    hjemme_score INT,
    borte_score INT,
    poeng INT,
    PRIMARY KEY(`bruker_id`, `kamp_id`),
    CONSTRAINT FK_Kamp_valgBruker FOREIGN KEY (bruker_id) REFERENCES bruker(bruker_id)
    --CONSTRAINT FK_Kamp_valgKamper FOREIGN KEY (kamp_id) REFERENCES kamper(kamp_id)
);

DROP TABLE IF EXISTS liganavn;
CREATE TABLE liganavn (
    liga_id INT NOT NULL AUTO_INCREMENT,
    liga_navn VARCHAR(128) NOT NULL,
    liga_kode VARCHAR(128) NOT NULL,
    PRIMARY KEY(`liga_id`),
    UNIQUE KEY unik_ligakode(`liga_kode`)
);

DROP TABLE IF EXISTS liga;
CREATE TABLE liga (
    liga_id INT NOT NULL,
    bruker_id INT NOT NULL,
    PRIMARY KEY(`liga_id`, `bruker_id`),
    CONSTRAINT FK_LigaBruker FOREIGN KEY (bruker_id) REFERENCES bruker(bruker_id),
    CONSTRAINT FK_LigaID FOREIGN KEY (liga_id) REFERENCES liganavn(liga_id)
);



DROP TABLE IF EXISTS medalje_valg;
CREATE TABLE medalje_valg (
    bruker_id INT NOT NULL,
    gull INT NOT NULL,
    solv INT NOT NULL,
    bronse INT NOT NULL,
    PRIMARY KEY(`bruker_id`),
    CONSTRAINT FK_Medalje_valgBruker FOREIGN KEY (bruker_id) REFERENCES bruker(bruker_id),
    CONSTRAINT FK_Gull_lag FOREIGN KEY (gull) REFERENCES lagnavn(lag_id),
    CONSTRAINT FK_Solv_lag FOREIGN KEY (solv) REFERENCES lagnavn(lag_id),
    CONSTRAINT FK_Bronse_lag FOREIGN KEY (bronse) REFERENCES lagnavn(lag_id)
);

--
-- Eksempel insert
--
INSERT INTO bruker
    (brukernavn, epost, passord)
VALUES
    ('Martin', 'martin@epost.no','12345'); -- spiller id blir automatisk generert