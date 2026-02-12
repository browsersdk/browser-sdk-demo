export interface AuthTokens {
  accessToken: string;
  refreshToken: string;
  expires: string;
  refreshExpire: string;
  username: string;
}

export class TokenManager {
  private static readonly ACCESS_TOKEN_KEY = 'access_token';
  private static readonly REFRESH_TOKEN_KEY = 'refresh_token';
  private static readonly EXPIRES_KEY = 'token_expires';
  private static readonly REFRESH_EXPIRES_KEY = 'refresh_expires';
  private static readonly USERNAME_KEY = 'username';

  static setTokens(tokens: AuthTokens): void {
    localStorage.setItem(this.ACCESS_TOKEN_KEY, tokens.accessToken);
    localStorage.setItem(this.REFRESH_TOKEN_KEY, tokens.refreshToken);
    localStorage.setItem(this.EXPIRES_KEY, tokens.expires);
    localStorage.setItem(this.REFRESH_EXPIRES_KEY, tokens.refreshExpire);
    localStorage.setItem(this.USERNAME_KEY, tokens.username);
  }

  static getAccessToken(): string | null {
    return localStorage.getItem(this.ACCESS_TOKEN_KEY);
  }

  static getRefreshToken(): string | null {
    return localStorage.getItem(this.REFRESH_TOKEN_KEY);
  }

  static getUsername(): string | null {
    return localStorage.getItem(this.USERNAME_KEY);
  }

  static isTokenExpired(): boolean {
    const expires = localStorage.getItem(this.EXPIRES_KEY);
    if (!expires) return true;
    
    return new Date(expires) <= new Date();
  }

  static isRefreshTokenExpired(): boolean {
    const refreshExpires = localStorage.getItem(this.REFRESH_EXPIRES_KEY);
    if (!refreshExpires) return true;
    
    return new Date(refreshExpires) <= new Date();
  }

  static clearTokens(): void {
    localStorage.removeItem(this.ACCESS_TOKEN_KEY);
    localStorage.removeItem(this.REFRESH_TOKEN_KEY);
    localStorage.removeItem(this.EXPIRES_KEY);
    localStorage.removeItem(this.REFRESH_EXPIRES_KEY);
    localStorage.removeItem(this.USERNAME_KEY);
  }

  static isAuthenticated(): boolean {
    return !!this.getAccessToken() && !this.isTokenExpired();
  }
}