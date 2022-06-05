# changefeed_cockroachDB

## Create Changefeed:-

                     CREATE CHANGEFEED FOR TABLE <table_name> INTO <sink_URL> ;

1. PAUSE A CHANGEFEED:-

                     Query:- PAUSE JOB <job_id>;

2.  RESUME A CHANGEFEED

                      Query:-  RESUME JOB <job_id>;

3.  CANCEL A CHANGEFEED:-

                     Query:- CANCEL JOB <job_id>;
