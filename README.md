# CDC_cockraochdb
Change Data capture incockroach db

Baseurl:  https://www.cockroachlabs.com/docs/stable/create-changefeed.html


2.	Creation Of ChangeFeed:


Example Query :-   CREATE CHANGEFEED FOR TABLE <table_name>  INTO  <sink_URL> ;


Query :-  CREATE CHANGEFEED FOR TABLE wallet INTO 'experimental- https://testfeed-twdwtabx5q-uc.a.run.app ' ;

1.	After executing this query ,It will create a Changefeed for the Table “wallet” with a specific JOB_ID.												
2.	To view the about of JOB_ID , The following Query will help to view the details of the JOB_ID.												
3.	Query:-  SELECT * FROM crdb_internal.jobs WHERE job_id= <job_id>;				
4.	To view about all JOBS created by ChangeFeed , The following Query will help to view.		
5.	Query:- SHOW JOBS;

NOTE-> Here we are using Google cloud run service(GCP) as sink to perform ChangeFeed.



3.	Operation With ChnageFeed:


1.	PAUSE A CHANGEFEED:-

Query:-  PAUSE JOB <job_id>;

2.	RESUME A CHANGEFEED

                      Query:-  RESUME JOB <job_id>;

3. CANCEL A CHANGEFEED:-

Query:- CANCEL JOB <job_id>;






For More Information :
Base URL: https://www.cockroachlabs.com/docs/stable/create-changefeed.html#manage-a-changefeed
