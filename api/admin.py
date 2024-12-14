from django.contrib import admin

from api.models import Author, Book, Genre, Editor
# Register your models here.

admin.site.register(Author)
admin.site.register(Book)
admin.site.register(Genre)
admin.site.register(Editor)
