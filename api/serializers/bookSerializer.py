from rest_framework import serializers
from api.models import Book
from api.models.Author import Author
from api.models.Genre import Genre
from api.models.Editor import Editor

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
            "authors",
            "editor",
            "genres"
        )

    def create(self, validated_data: any) -> Book:
        return Book.objects.create(**validated_data)
    
    def update(self, instance: Book, validated_data: any) -> Book:
        instance.name = validated_data.get("name")
        instance.parution = validated_data.get("parution")
        instance.price = validated_data.get("price")
        instance.synopsis = validated_data.get("synopsis")
        instance.ISBN = validated_data.get("ISBN")
        instance.stock = validated_data.get("stock")
        instance.new = validated_data.get("new")
        instance.note = validated_data.get("note")
        # TODO: define the elements already in, the elements that need to be removed and the elements to be added
        instance.authors.clear()
        for author_id in validated_data.get("authors"):
            instance.authors.add(
                Author.objects.get(pk=author_id)
            )
        if instance.editor.pk != validated_data.get("editor"):
            instance.editor = Editor.objects.get(pk=validated_data.get("editor"))
        # TODO same as authors
        instance.genres.clear()
        for type_id in validated_data.get("genres"):
            instance.genres.add(
                Genre.objects.get(pk=type_id)
            )
        instance.save()
        return instance