import { HTTPService } from './httpService';
import { AuthService } from './authService';

export const httpService = new HTTPService();
export const authService = new AuthService(httpService);
