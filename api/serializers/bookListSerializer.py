from rest_framework import serializers
from api.models import BookList

class BookListSerializer(serializers.ModelSerializer):
    class Meta:
        model = BookList
        fields = (
            "id",
            "name",
            "description",
            "priority",
            "books"
        )
    
    def create(self, validated_data: any) -> BookList:
        bookList = BookList(
            name        = validated_data.get("name"),
            description = validated_data.get("description"),
            priority    = validated_data.get("priority")
        )
        bookList.save()
        for book in validated_data.get("books"):
            bookList.books.add(book)
        bookList.save()
        return bookList
    
    def update(self, instance: BookList, validated_data: any) -> BookList:
        instance.name        = validated_data.get("name")
        instance.description = validated_data.get("description")
        instance.priority    = validated_data.get("priority")
        
        instance.books.clear()
        for book in validated_data.get("books"):
            instance.books.add(book)
        instance.save()
        return instance