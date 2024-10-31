import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ApiResponse } from '../models/apiresponse.model';
import { ITask } from '../models/tasks.model';

@Injectable({
  providedIn: 'root'
})
export class TaskService {
  apiUrl = 'http://localhost:8089/tasks/';

  constructor(private http: HttpClient) {}

  // GET all tasks
  getAllTasks(): Observable<ApiResponse<ITask[]>> {
    return this.http.get<ApiResponse<ITask[]>>(this.apiUrl);
  }

  // GET a specific task
  getTask(id: number): Observable<ApiResponse<ITask>> {
    return this.http.get<ApiResponse<ITask>>(`${this.apiUrl}${id}`);
  }

  // POST a new task
  addTask(task: ITask): Observable<ApiResponse<ITask>> {
    return this.http.post<ApiResponse<ITask>>(this.apiUrl, task);
  }

  // PUT to update an existing task
  updateTask(id: number, task: ITask): Observable<ApiResponse<ITask>> {
    return this.http.put<ApiResponse<ITask>>(`${this.apiUrl}${id}`, task);
  }

  // DELETE a task
  deleteTask(id: number): Observable<any> {
    return this.http.delete(`${this.apiUrl}${id}`);
  }
}
