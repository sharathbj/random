from django.shortcuts import render
from .models import Todo
from django.http import HttpResponseRedirect

def todoView(request):
	all_todo_items = Todo.objects.all()
	return render(request,'home.html',{'all_items':all_todo_items})

def addTask(request):
	new_item = Todo(text = request.POST.get('task','false'))
	new_item.save()
	return HttpResponseRedirect('/todo/')

def deleteTask(request,todo_id):
        del_item = Todo.objects.get(id=todo_id)
        del_item.delete()
        return HttpResponseRedirect('/todo/')	
	
