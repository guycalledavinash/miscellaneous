export type Player={id:number;first_name:string;last_name:string;nationality:string;club:string;position:string;age:number;goals:number;assists:number;market_value:number};
export type Team={id:number;name:string;country:string;coach:string;fifa_ranking:number};
export type Match={id:number;home_team_id:number;away_team_id:number;stadium:string;kickoff_time:string;home_score:number|null;away_score:number|null;competition_id:number};
export type Dashboard={total_players:number;total_teams:number;upcoming_matches:number;top_scorers:Player[];top_assists:Player[];team_rankings:Team[];nationality_distribution:Record<string,number>};
