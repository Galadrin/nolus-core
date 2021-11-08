import {SigningCosmWasmClient} from "@cosmjs/cosmwasm-stargate";
import {DirectSecp256k1Wallet} from "@cosmjs/proto-signing";
import {fromHex} from "@cosmjs/encoding";

let validatorPrivKey = fromHex(process.env.VALIDATOR_PRIV_KEY as string);
let periodicPrivKey = fromHex(process.env.PERIODIC_PRIV_KEY as string);

export async function getWallet(privKey: Uint8Array): Promise<DirectSecp256k1Wallet> {
    return await DirectSecp256k1Wallet.fromKey(privKey, "nomo");
}

export async function getClient(privKey: Uint8Array): Promise<SigningCosmWasmClient> {
    return await SigningCosmWasmClient.connectWithSigner(process.env.NODE_URL as string, await getWallet(privKey));
}

export async function getValidatorWallet(): Promise<DirectSecp256k1Wallet> {
    return await getWallet(validatorPrivKey);
}

export async function getValidatorClient(): Promise<SigningCosmWasmClient> {
    return await getClient(validatorPrivKey);
}

export async function getPeriodicWallet(): Promise<DirectSecp256k1Wallet> {
    return await getWallet(periodicPrivKey);
}

export async function getPeriodicClient(): Promise<SigningCosmWasmClient> {
    return await getClient(periodicPrivKey);
}
