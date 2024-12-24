from rest_framework import serializers
from api.models import Book

class BookSerializer(serializers.ModelSerializer):
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
        book = Book(
            name     = validated_data.get("name"),
            parution = validated_data.get("parution"),
            price    = validated_data.get("price"),
            synopsis = validated_data.get("synopsis"),
            ISBN     = validated_data.get("ISBN"),
            stock    = validated_data.get("stock"),
            new      = validated_data.get("new"),
            note     = validated_data.get("note"),
            editor   = validated_data.get("editor")
        )
        book.save()
        for auth in validated_data.get("authors"):
            book.authors.add(auth)
        for genre in validated_data.get("genres"):
            book.genres.add(genre)
        book.save()        
        return book
    
    def update(self, instance: Book, validated_data: any) -> Book:
        instance.name     = validated_data.get("name")
        instance.parution = validated_data.get("parution")
        instance.price    = validated_data.get("price")
        instance.synopsis = validated_data.get("synopsis")
        instance.ISBN     = validated_data.get("ISBN")
        instance.stock    = validated_data.get("stock")
        instance.new      = validated_data.get("new")
        instance.note     = validated_data.get("note")
        # TODO: define the elements already in, the elements that need to be removed and the elements to be added
        instance.authors.clear()
        for author in validated_data.get("authors"):
            instance.authors.add(author)
        if instance.editor.pk != validated_data.get("editor"):
            instance.editor = validated_data.get("editor")
        # TODO same as authors
        instance.genres.clear()
        for genre in validated_data.get("genres"):
            instance.genres.add(genre)
        instance.save()
        return instance