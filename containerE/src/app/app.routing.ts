import { Routes, RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { DashboardComponent } from '@pages/dashboard/dashboard.component';
import { CharactersComponent } from '@pages/dashboard/characters/characters.component';
import { EpisodesComponent } from '@pages/dashboard/episodes/episodes.component';
import { LocationsComponent } from '@pages/dashboard/locations/locations.component';
import { HomeComponent } from '@pages/dashboard/home/home.component';

const routes: Routes = [
    {
        path: '', component: DashboardComponent, children: [
            { path: '', pathMatch: 'full', redirectTo: 'home' },
            { path: 'home', component: HomeComponent },
            { path: 'characters', component: CharactersComponent },
            { path: 'episodes', component: EpisodesComponent },
            { path: 'locations', component: LocationsComponent },
        ]
    },
    { path: '**', pathMatch: 'full', redirectTo: 'home' }
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule { }
