from rest_framework import serializers
from api.models import Book

class BookListSerializer(serializers.ModelSerializer):
    class Meta:
        model = Book
        fields = (
            "id",
            "name",
            "parution",
            "price",
            "synopsis",
            "ISBN",
            "stock",
            "new",
            "note",
            "author",
            "editor",
            "types"
        )

    def create(self, validated_data: any) -> Book:
        return Book.objects.create(**validated_data)
    
    def update(self, instance: Book, validated_data: any) -> Book:
        instance.name = validated_data.get("name")
        instance.surname = validated_data.get("surname")
        instance.save()
        return instance