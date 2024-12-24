from rest_framework import serializers
from api.models import Library

class LibrarySerializer(serializers.ModelSerializer):
    class Meta:
        model = Library
        fields = (
            "id",
            "name",
            "town",
            "address_1",
            "address_2",
            "phone",
            "email",
            "about"
        )
    
    def create(self, validated_data: any) -> Library:
        lib = Library(
            name=validated_data.get("name"),
            town=validated_data.get("town"),
            address_1=validated_data.get("address_1"),
            address_2=validated_data.get("address_2"),
            phone=validated_data.get("phone"),
            email=validated_data.get("email"),
            about=validated_data.get("about")
        )
        lib.save()
        return lib
    
    def update(self, instance: Library, validated_data: any) -> Library:
        instance.name=validated_data.get("name")
        instance.town=validated_data.get("town")
        instance.address_1=validated_data.get("address_1")
        instance.address_2=validated_data.get("address_2")
        instance.phone=validated_data.get("phone")
        instance.email=validated_data.get("email")
        instance.about=validated_data.get("about")
        instance.save()
        return instance