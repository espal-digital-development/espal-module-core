import { HTTPService } from './http';
import { AuthService } from './auth';
import { StorageService } from './storage';

export const httpService = new HTTPService();
export const storageService = new StorageService()
export const authService = new AuthService(httpService, storageService);
