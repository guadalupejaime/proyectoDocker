import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Characters } from '@pages/dashboard/characters/characters.interface';
import { TablePage } from '@services/http.interfaces';
import { urlList } from '@services/urlList';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  constructor(
    private sHttpClient: HttpClient
  ) { }

  private makeUrlParams(perametros: {}): string {
    let temp = '';
    for (const key in perametros) {
      if (perametros.hasOwnProperty(key)) {
        const element = perametros[key];
        if (element || element === 0) {
          temp += key + '=' + element + '&';
          // temp += key + '=' + encodeURIComponent(element) + '&';
        }
      }
    }
    if (temp.length) {
      temp = temp.slice(0, temp.length - 1);
    }
    return temp;
  }

  getCharacters(pagina: TablePage = { offset: 0, limit: 100 }): Observable<Characters> {
    const myUrl = urlList.characters + '?' + this.makeUrlParams(pagina);
    return this.sHttpClient.get<Characters>(myUrl, {});
  }
}
