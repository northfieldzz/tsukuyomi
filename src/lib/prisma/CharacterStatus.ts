import {CharacterStatus} from "@prisma/client";

class Dice {
    sides: number

    constructor(sides: number) {
        this.sides = sides
    }

    roll(count: number) {
        let result = 0;
        for (let i = 0; i < count; i++) {
            result += Math.floor(Math.random() * this.sides) + 1
        }
        return result
    }
}

export class CharacterGenerator {
    app: number
    con: number
    dex: number
    edu: number
    int: number
    luk: number
    pow: number
    siz: number
    str: number
    san: number
    ida: number
    know: number
    hp: number
    mov: number
    mp: number
    skl: number
    db: number
    build: number

    public static generate() {
        const dice = new Dice(6)
        const pow = dice.roll(3) * 5
        const int = (dice.roll(2) + 6) * 5
        const edu = (dice.roll(2) + 6) * 5
        const con = dice.roll(3) * 5
        const siz = (dice.roll(2) + 6) * 5
        const str = dice.roll(3) * 5
        const dex = dice.roll(3) * 5
        let mov = 7
        if (dex < siz && dex < str) {
            mov = 7
        } else if ((dex >= siz || dex >= siz) || (dex === str && siz === str)) {
            mov = 8
        } else {

        }
        return new this({
            app: dice.roll(3) * 5,
            con: con,
            dex: dex,
            edu: edu,
            int: int,
            luk: dice.roll(3),
            pow: pow,
            siz: siz,
            str: str,
            san: pow,
            ida: int,
            know: edu,
            hp: (con + siz) / 10,
            mov: mov,
            mp: pow / 5,
            skl: int * 2,
            db: str + siz,
            build: str + siz
        })
    }

    constructor(data: Omit<CharacterStatus, 'id' | 'guildId' | 'userId'>) {
        this.app = data.app
        this.con = data.con
        this.dex = data.dex
        this.edu = data.edu
        this.int = data.int
        this.luk = data.luk
        this.pow = data.pow
        this.siz = data.siz
        this.str = data.str
        this.luk = data.luk
        this.san = data.san
        this.ida = data.ida
        this.know = data.know
        this.hp = data.hp
        this.mov = data.mov
        this.mp = data.mp
        this.skl = data.skl
        this.db = data.db
        this.build = data.build
    }
}

function trans(value: number) {
    const valueStr = value.toString()
    let regex: RegExp
    switch (valueStr.length) {
        case 1:
            regex = /^0+/
            return valueStr.padStart(3, '0').replace(regex, '  ')
        case 2:
            regex = /^0/
            return valueStr.padStart(3, '0').replace(regex, ' ')
        case 3:
        default:
            return valueStr
    }
}

export function summary(username: string | undefined | null, status: Omit<CharacterStatus, 'id' | 'guildId' | 'userId'>) {
    return `
\`\`\`txt
${username} Status
|APP:   ${trans(status.app)} |CON  :   ${trans(status.con)} |
|DEX:   ${trans(status.dex)} |EDU  :   ${trans(status.edu)} |
|INT:   ${trans(status.int)} |LUK  :   ${trans(status.luk)} |
|POW:   ${trans(status.pow)} |SIZ  :   ${trans(status.siz)} |
|STR:   ${trans(status.str)} |SAN  :   ${trans(status.san)} |
|IDA:   ${trans(status.ida)} |KNOW :   ${trans(status.know)} |
|HP :   ${trans(status.hp)} |MOV  :   ${trans(status.mov)} |
|MP :   ${trans(status.mp)} |SKL  :   ${trans(status.skl)} |
|DB :   ${trans(status.db)} |BUILD:   ${trans(status.build)} |
\`\`\``
}