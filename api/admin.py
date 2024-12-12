from django.contrib import admin

from api.models import Author, Book, BookType, Editor
# Register your models here.

admin.site.register(Author)
admin.site.register(Book)
admin.site.register(BookType)
admin.site.register(Editor)
