from rest_framework import serializers
from api.models import Genre

class GenreSerializer(serializers.ModelSerializer):
    class Meta:
        model = Genre
        fields = (
            "id",
            "name"
        )
    def create(self, validated_data: any) -> Genre:
        return Genre.objects.create(**validated_data)
    
    def update(self, instance: Genre, validated_data: any) -> Genre:
        instance.name = validated_data.get("name")
        instance.save()
        return instance
