create table user_info(
   id           text                not null,
   name         text                not null,
   password     text                not null,
   phonenumber  text   primary key  not null,
   admin       boolean             not null,
   integral     int                 not null,
   user_image   text
);

create table seat_info(
   id           text    primary key not null,
   fenguan      text                not null,
   louceng      text                not null,
   bianhao      text                not null,
   number       int                 not null,
   status       int                 not null,
   begintime    text,
   endtime      text,
   date         text,
   seat_image   text
);

-- 座位信息需要将座位的所有信息进行录入
create table seat_info_detailed(
   id           text    primary key not null,
   fenguan      text                not null,
   louceng      text                not null,
   bianhao      text                not null,
   number       int                 not null,
   status       int                 not null,
   begintime    text,
   endtime      text,
   date         text,
   seat_image   text,
   -- 前排后排 设定门开始为第一排
   line         text                not null,
   -- 是否靠窗
   nearwindow   boolean             not null,
   -- 是否向阳
   tosum        boolean             not null,
   -- 是否背光
   backsum      boolean             not null,
   -- 是否有空调
   airconditioner boolean           not null,
   -- 是否是单人座位(没有同桌)
   singleseat   boolean             not null,
   -- 是否多人座位(有至少一个同桌)
   multipleseat boolean             not null
);

create table user_his_detail_info(
   id           text    primary key not null,
   phonenumber  text                not null,
   fenguan      text                not null,
   louceng      text                not null,
   bianhao      text                not null,
   number       int                 not null,
   begintime    text                not null,
   endtime      text                not null,
   date         text                not null,
   status       int                not null,

   -- 前排后排 设定门开始为第一排
   line         text                not null,
   -- 是否靠窗
   nearwindow   boolean             not null,
   -- 是否向阳
   tosum        boolean             not null,
   -- 是否背光
   backsum      boolean             not null,
   -- 是否有空调
   airconditioner boolean           not null,
   -- 是否是单人座位(没有同桌)
   singleseat   boolean             not null,
   -- 是否多人座位(有至少一个同桌)
   multipleseat boolean             not null
);

create table user_hisinfo(
   id           text    primary key not null,
   phonenumber  text                not null,
   fenguan      text                not null,
   louceng      text                not null,
   bianhao      text                not null,
   number       int                 not null,
   begintime    text                not null,
   endtime      text                not null,
   date         text                not null,
   status       int                not null
);

create table user_re_back(
   id           text    primary key not null,
   userphone      text        not null,
   usermsgtext    text        not null,
   umsgimage      text        not null,
   adminphone     text        not null,
   adminmsgtext   text        not null,
   amsgimage      text        not null,
   status         boolean     not null
)