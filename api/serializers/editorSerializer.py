from rest_framework import serializers
from api.models import Editor

class EditorSerializer(serializers.ModelSerializer):
    class Meta:
        model = Editor
        fields = (
            "id",
            "name"
        )
    def create(self, validated_data: any) -> Editor:
        return Editor.objects.create(**validated_data)
    
    def update(self, instance: Editor, validated_data: any) -> Editor:
        instance.name = validated_data.get("name")
        instance.save()
        return instance
