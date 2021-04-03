export interface Episode {
    air_date: string;
    characters: string[];
    created: string;
    episode: string;
    id: number;
    name: string;
    checked?: boolean;
}

export interface Info {
    count: number;
    next: string;
    page: number;
    prev: string;
}

export interface Episodes {
    info: Info;
    results: Episode[];
    total_found: number;
    total_returned: number;
}
