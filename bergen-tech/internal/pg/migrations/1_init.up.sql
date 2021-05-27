CREATE TABLE "job" (
	"ID" serial NOT NULL,
	"Name" varchar(255) NOT NULL UNIQUE,
	CONSTRAINT "job_pk" PRIMARY KEY ("ID")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "task" (
	"ID" serial NOT NULL,
	"Name" varchar(255) NOT NULL,
	"jobID" bigint NOT NULL,
	CONSTRAINT "task_pk" PRIMARY KEY ("ID")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "labor" (
	"ID" serial NOT NULL,
	"Name" varchar(255) NOT NULL,
	"taskID" bigint NOT NULL,
	"PayTime" bigint NOT NULL,
	CONSTRAINT "labor_pk" PRIMARY KEY ("ID")
) WITH (
  OIDS=FALSE
);

ALTER TABLE "task"  ADD CONSTRAINT "task_fk0"  FOREIGN KEY ("jobID")  REFERENCES "job"("ID");
ALTER TABLE "labor" ADD CONSTRAINT "labor_fk0" FOREIGN KEY ("taskID") REFERENCES "task"("ID");
