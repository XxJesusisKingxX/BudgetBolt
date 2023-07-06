PGDMP     /    -    	            {        
   budgetbolt    15.3    15.3 "               0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                        1262    16398 
   budgetbolt    DATABASE     �   CREATE DATABASE budgetbolt WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.1252';
    DROP DATABASE budgetbolt;
                postgres    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                pg_database_owner    false            !           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                   pg_database_owner    false    4            �            1259    16405    budget    TABLE     �   CREATE TABLE public.budget (
    budget_id integer NOT NULL,
    budget_name character(25) NOT NULL,
    short_description character(50),
    budget_frequency character(20) NOT NULL
);
    DROP TABLE public.budget;
       public         heap    postgres    false    4            �            1259    16425    budget_budget_id_seq    SEQUENCE     �   ALTER TABLE public.budget ALTER COLUMN budget_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.budget_budget_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    4    214            �            1259    16431    expense    TABLE       CREATE TABLE public.expense (
    expense_id integer NOT NULL,
    due_date date,
    expense_name character(25) NOT NULL,
    expense_limit double precision,
    budget_id integer NOT NULL,
    transaction_id integer,
    expense_category character(100)
);
    DROP TABLE public.expense;
       public         heap    postgres    false    4            �            1259    16435    expense_expense_id_seq    SEQUENCE     �   ALTER TABLE public.expense ALTER COLUMN expense_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.expense_expense_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    4    218            �            1259    16451    income    TABLE     �   CREATE TABLE public.income (
    income_id integer NOT NULL,
    income_name character(25) NOT NULL,
    income_amount_expected double precision NOT NULL,
    budget_id integer NOT NULL,
    transaction_id integer,
    income_category character(100)
);
    DROP TABLE public.income;
       public         heap    postgres    false    4            �            1259    16450    income_income_id_seq    SEQUENCE     �   ALTER TABLE public.income ALTER COLUMN income_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.income_income_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    4    221            �            1259    16419    transaction    TABLE     f  CREATE TABLE public.transaction (
    transaction_id integer NOT NULL,
    transaction_date date NOT NULL,
    net_amount double precision NOT NULL,
    payment_method character(20) NOT NULL,
    payment_account_from_to character(50) NOT NULL,
    vendor character(50),
    is_recurring boolean DEFAULT false NOT NULL,
    short_description character(50)
);
    DROP TABLE public.transaction;
       public         heap    postgres    false    4            �            1259    16418    transaction_transaction_id_seq    SEQUENCE     �   ALTER TABLE public.transaction ALTER COLUMN transaction_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.transaction_transaction_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    216    4                      0    16405    budget 
   TABLE DATA           ]   COPY public.budget (budget_id, budget_name, short_description, budget_frequency) FROM stdin;
    public          postgres    false    214   �(                 0    16431    expense 
   TABLE DATA           �   COPY public.expense (expense_id, due_date, expense_name, expense_limit, budget_id, transaction_id, expense_category) FROM stdin;
    public          postgres    false    218   �(                 0    16451    income 
   TABLE DATA           |   COPY public.income (income_id, income_name, income_amount_expected, budget_id, transaction_id, income_category) FROM stdin;
    public          postgres    false    221   �(                 0    16419    transaction 
   TABLE DATA           �   COPY public.transaction (transaction_id, transaction_date, net_amount, payment_method, payment_account_from_to, vendor, is_recurring, short_description) FROM stdin;
    public          postgres    false    216   �(       "           0    0    budget_budget_id_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.budget_budget_id_seq', 1, false);
          public          postgres    false    217            #           0    0    expense_expense_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public.expense_expense_id_seq', 1, false);
          public          postgres    false    219            $           0    0    income_income_id_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.income_income_id_seq', 1, false);
          public          postgres    false    220            %           0    0    transaction_transaction_id_seq    SEQUENCE SET     M   SELECT pg_catalog.setval('public.transaction_transaction_id_seq', 1, false);
          public          postgres    false    215            v           2606    16430    budget budget_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.budget
    ADD CONSTRAINT budget_pkey PRIMARY KEY (budget_id);
 <   ALTER TABLE ONLY public.budget DROP CONSTRAINT budget_pkey;
       public            postgres    false    214            z           2606    16437    expense expense_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.expense
    ADD CONSTRAINT expense_pkey PRIMARY KEY (expense_id);
 >   ALTER TABLE ONLY public.expense DROP CONSTRAINT expense_pkey;
       public            postgres    false    218            �           2606    16455    income income_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.income
    ADD CONSTRAINT income_pkey PRIMARY KEY (income_id);
 <   ALTER TABLE ONLY public.income DROP CONSTRAINT income_pkey;
       public            postgres    false    221            x           2606    16423    transaction transaction_pkey 
   CONSTRAINT     f   ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_pkey PRIMARY KEY (transaction_id);
 F   ALTER TABLE ONLY public.transaction DROP CONSTRAINT transaction_pkey;
       public            postgres    false    216            {           1259    16449    fki_fk_expense_budget_id    INDEX     Q   CREATE INDEX fki_fk_expense_budget_id ON public.expense USING btree (budget_id);
 ,   DROP INDEX public.fki_fk_expense_budget_id;
       public            postgres    false    218            |           1259    16443    fki_fk_expense_transaction_id    INDEX     [   CREATE INDEX fki_fk_expense_transaction_id ON public.expense USING btree (transaction_id);
 1   DROP INDEX public.fki_fk_expense_transaction_id;
       public            postgres    false    218            }           1259    16461    fki_fk_income_budget_id    INDEX     O   CREATE INDEX fki_fk_income_budget_id ON public.income USING btree (budget_id);
 +   DROP INDEX public.fki_fk_income_budget_id;
       public            postgres    false    221            ~           1259    16467    fki_fk_income_transaction_id    INDEX     Y   CREATE INDEX fki_fk_income_transaction_id ON public.income USING btree (transaction_id);
 0   DROP INDEX public.fki_fk_income_transaction_id;
       public            postgres    false    221            �           2606    16444    expense fk_expense_budget_id    FK CONSTRAINT     �   ALTER TABLE ONLY public.expense
    ADD CONSTRAINT fk_expense_budget_id FOREIGN KEY (budget_id) REFERENCES public.budget(budget_id);
 F   ALTER TABLE ONLY public.expense DROP CONSTRAINT fk_expense_budget_id;
       public          postgres    false    3190    218    214            �           2606    16438 !   expense fk_expense_transaction_id    FK CONSTRAINT     �   ALTER TABLE ONLY public.expense
    ADD CONSTRAINT fk_expense_transaction_id FOREIGN KEY (transaction_id) REFERENCES public.transaction(transaction_id);
 K   ALTER TABLE ONLY public.expense DROP CONSTRAINT fk_expense_transaction_id;
       public          postgres    false    218    3192    216            �           2606    16456    income fk_income_budget_id    FK CONSTRAINT     �   ALTER TABLE ONLY public.income
    ADD CONSTRAINT fk_income_budget_id FOREIGN KEY (budget_id) REFERENCES public.budget(budget_id);
 D   ALTER TABLE ONLY public.income DROP CONSTRAINT fk_income_budget_id;
       public          postgres    false    214    221    3190            �           2606    16462    income fk_income_transaction_id    FK CONSTRAINT     �   ALTER TABLE ONLY public.income
    ADD CONSTRAINT fk_income_transaction_id FOREIGN KEY (transaction_id) REFERENCES public.transaction(transaction_id);
 I   ALTER TABLE ONLY public.income DROP CONSTRAINT fk_income_transaction_id;
       public          postgres    false    216    3192    221                  x������ � �            x������ � �            x������ � �            x������ � �     