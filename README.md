# Lies of P Essential Mods

I love Lies of P, however it has some shit balancing. This git repo puts some mods together and provides a launcher for Linux.

Oh yeah make sure to pick the right P-Organ upgrades cause some of them should be there by default.

-   **No Random Damage**<br />
    The game does `+- 10% damage`. No other souls game does. Hits should be consistent. Test this at Hotel Krat on a dummy.<br />
    https://www.nexusmods.com/liesofp/mods/175

-   **Improved Perfect Guard**<br />
    Timing from `155 ms` to `200 ms`<br />
    Recovery from `500 ms` to `350 ms`<br />
    Sekiro's timing is `500 ms` lmao. This game is way too tight.<br />
    There's an extended `250 ms` version below too.<br />
    https://www.nexusmods.com/liesofp/mods/161

-   **No Intro Skip**<br />
    https://www.nexusmods.com/liesofp/mods/27

Just merge this repo into the game's files and set Steam launch args to:

```bash
gamemoderun ./NoRandomDamageLauncher %command%
```
