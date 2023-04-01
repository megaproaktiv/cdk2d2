import { Construct } from 'constructs';


export function  Show(c: Construct) {
        c.node.addMetadata("Show", "true")
}

export function  Connection(source: Construct, target: Construct) {
    source.node.addMetadata("Connection", target.node.id)
}
export function  Container(source: Construct, target: Construct) {
    source.node.addMetadata("Container", target.node.id)
}
