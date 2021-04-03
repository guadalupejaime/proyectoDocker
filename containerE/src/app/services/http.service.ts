import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Character, Characters } from '@pages/dashboard/characters/characters.interface';
import { Episode, Episodes } from '@pages/dashboard/episodes/episodes.interface';
import { Locations } from '@pages/dashboard/locations/locations.interface';
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

  /** characters */
  getCharacters(pagina: TablePage = { offset: 0, limit: 100 }): Observable<Characters> {
    const myUrl = urlList.characters + '?' + this.makeUrlParams(pagina);
    return this.sHttpClient.get<Characters>(myUrl);
  }

  getCharacter(idCharacter: string): Observable<Character> {
    const myUrl = urlList.characters + '/' + idCharacter;
    return this.sHttpClient.get<Character>(myUrl);
  }

  postCharacters(character: Character): any {
    const myUrl = urlList.characters + '/rabbit';
    return this.sHttpClient.post(myUrl, character);
  }

  /** episodes */
  getEpisodes(pagina: TablePage = { offset: 0, limit: 108 }): Observable<Episodes> {
    const myUrl = urlList.episodes + '?' + this.makeUrlParams(pagina);
    return this.sHttpClient.get<Episodes>(myUrl);
  }

  getEpisode(idEpisode: string): Observable<Episode> {
    const myUrl = urlList.episodes + '/' + idEpisode
    return this.sHttpClient.get<Episode>(myUrl);
  }

  postEpisode(episode: Episode): any {
    const myUrl = urlList.episodes + '/rabbit';
    return this.sHttpClient.post(myUrl, episode);
  }

  /** locations */
  getLocations(pagina: TablePage = { offset: 0, limit: 100 }): Observable<Locations> {
    const myUrl = urlList.locations + '?' + this.makeUrlParams(pagina);
    return this.sHttpClient.get<Locations>(myUrl);
  }
}
