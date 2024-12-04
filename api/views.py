from django.http import JsonResponse
from django.shortcuts import render

from api.models import Author
from api.serializers import AuthorSerializer

from django.views.decorators.csrf import csrf_exempt

# Create your views here.

@csrf_exempt
def get_data(request):
    data = Author.objects.all()
    if request.method == "GET":
        serializer = AuthorSerializer(data, many=True)
        return JsonResponse(serializer.data, safe=False)
