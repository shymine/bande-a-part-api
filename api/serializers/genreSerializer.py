from rest_framework import serializers
from api.models import Author

class GenreSerializer(serializers.ModelSerializer):
    class Meta:
        model = Author
        fields = (
            "id",
            "name",
            "surname"
        )
    def create(self, validated_data: any) -> Author:
        return Author.objects.create(**validated_data)
    
    def update(self, instance: Author, validated_data: any) -> Author:
        instance.name = validated_data.get("name")
        instance.surname = validated_data.get("surname")
        instance.save()
        return instance
