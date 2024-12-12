# Commands

## Database

- when making modifications (or creating new) on models: 
    - `python manage.py makemigrations {app name}`
        - create the database commit
    - `python manage.py sqlmigrate {app name} {commit version}`
        -  show the comit content
    - `python manage.py migrate`
        - apply the migration

- create superuser:
    - `python manage.py createsuperuser`

- reseting the db:
    - `python manage.py reset_db`
        - clear the tables
    - `python manage.py syncdb`
        - will recreate them

## Operations point

- doing punctual addition
    - `python manage.py shell`