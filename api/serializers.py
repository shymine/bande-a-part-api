from rest_framework import serializers
from api.models import Author, BookType, Editor, Book

class AuthorSerializer(serializers.ModelSerializer):
    class Meta:
        model = Author
        fields = (
            "name",
            "surname"
        )

