from api.models import Editor
from api.serializers.editorSerializer import EditorSerializer
from rest_framework import generics

class EditorListView(generics.ListCreateAPIView):
    queryset = Editor.objects.all()
    serializer_class = EditorSerializer

class EditorDetail(generics.RetrieveUpdateDestroyAPIView):
    queryset = Editor.objects.all()
    serializer_class = EditorSerializer