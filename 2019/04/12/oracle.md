
https://groups.google.com/forum/#!topic/jlancer/xLUKfwGnjZo


오라클 RAC환경에서 JDBC 접속 URL 설정

 

 

오라클의 RAC(Real Appliction Clusters)은 오라클사의 데이터베이스 클러스터링 솔루션으로

여러대의 DB를 동일한 서비스 네임으로 접속 하여 마치 하나의 DB에 접속하여 사용하는 것처럼

구성해 주는 서비스다. 따라서 어떤 DB에 연결되어지는 자동으로 연결되어 지며 한곳에 등록된

데이터는 다른 DB에도 동기화 처리 된다.

 

SID1, SID2 의 두대의 서버를 RAC로 구성 한다면 두서버를 하나의 서비스네임인 SID로 구성 하여

사용들은 SID로 접속하면 내부적으로 SID1 또는 SID2 로 분사처리 되면 한쪽 DB SID1가 FailOver나면

자동으로 다른 DB SID2로 연결 된다.

 

개발시 유의 사항

 

- 업무 파티션을 하고자 할 경우 각 업무별로 장애가 발생하지 않는 평시에
  RACDB1으로 접속하여 사용할 것인지 RACDB2로 접속하여 사용할 것인지를 결정하여야 함.
- 업무파티션의 기준은 각각의 노드에서 테이블 단위의 DML(INSERT,UPDATE,DELETE)이
  최소화되도록 설계하여야 함.

 

JDBC를 통해 RAC로 구성된 오라클 DB에 접속 할 때는 Single DB를 사용 할 때와 차이가 난다.

FailOver 와 LoadBlance 등과 각 DB TNS 정보를 설정하여 주어야 한다.

 

Single DB: jdbc:oracle:thin:@xxx.xxx.xxx:1521:SID

 

RAC DB:  jdbc:oracle:thin:@(DESCRIPTION=
                           (FAIL_OVER=ON)
                           (LOAD_BALANCE=ON
                          (ADDRESS=(PROTOCOL=TCP)(HOST=xxx.xxx.xxx.xxx)

                                    (PORT=1521)) # SID1 DB
                           (ADDRESS=(PROTOCOL=TCP)(HOST=xxx.xxx.xxx.xxx)

                                    (PORT=1521)) # SID2 DB
                           (CONNECT_DATA=(SERVICE_NAME=SID)))

 

* 옵션: FAIL_OVER=ON (장애발생시 연결서버 변경)

       LOAD_BALANCE=ON (분산처리)

 

-------------------------------------------------------------------------------------

 

* 오라클 리스너 설정(listener.ora)

 

   # Local Listener를 사용하도록 설정

   LISTENER_HOSTNAME =
     (DESCRIPTION_LIST =
       (DESCRIPTION =
         (ADDRESS = (PROTOCOL = TCP)(HOST = VIP_HOSTNAME)(PORT = 1521)(IP=FIRST))
         (ADDRESS = (PROTOCOL = TCP)(HOST = HOSTNAME)    (PORT = 1521)(IP=FIRST))
         (ADDRESS = (PROTOCOL = IPC)(KEY = extproc))
       )
     )
 

   SID_LIST_LISTENER_HOSTNAME =
     (SID_LIST =
       (SID_DESC =
         (SID_NAME = PLSExtProc)
         (ORACLE_HOME = /oracle/app/oracle/product/10.2.0/db_1)
         (PROGRAM = extproc)
      

 

* TNS(tnsnames.ora) 설정 필요 - JDBC OCI 방식

    # JDBC OCI 방식으로 설정하여 사용하고자 할 경우 TNS설정이 필요하며
    # RACDB에 접속하는 방식은 SQL*Net을 사용하게 됨
    # (SELECT인 것만 FAILOVER할 수 있도록 설정한 예, 주로 많이 사용하는 설정)

   TNS_RACDB1 = (DESCRIPTION=
                 (LOAD_BALANCE=OFF)
                 (FAILOVER=ON)
                 (ADDRESS=(PROTOCOL=TCP)(HOST=VIP_HOSTNAME1)(PORT=1521))
                 (ADDRESS=(PROTOCOL=TCP)(HOST=VIP_HOSTNAME2)(PORT=1521))
                 (CONNECT_DATA = (SERVER = DEDICATED)
                                 (SERVICE_NAME = RACDB)
                                 (FAILOVER_MODE=(TYPE=SELECT)(METHOD=BASIC))
                 )
             )
   
   TNS_RACDB2 = (DESCRIPTION=
                 (LOAD_BALANCE=OFF)
                 (FAILOVER=ON)
                 (ADDRESS=(PROTOCOL=TCP)(HOST=VIP_HOSTNAME2)(PORT=1521))
                 (ADDRESS=(PROTOCOL=TCP)(HOST=VIP_HOSTNAME1)(PORT=1521))
                 (CONNECT_DATA = (SERVER = DEDICATED)
                                 (SERVICE_NAME = RACDB)
                                 (FAILOVER_MODE=(TYPE=SELECT)(METHOD=BASIC))
                 )
             )

 

     1]  RACDB1에 접속하고자 하는 설정
          jdbc:oracle:oci:@TNS_RACDB1

 

     2  RACDB2 접속하고자 하는 설정
          jdbc:oracle:oci:@TNS_RACDB2